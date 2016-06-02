# -*- coding: utf-8 -*-

__author__ = 'wills'

import os
import sys
import getopt
import yaml
import re

class Spec(object):

    DB_TYPE_TO_PROTO_MAPPING = {
        "integer": "int32",
        "smallint": "int32",
        "string": "string",
        "float": "float",
        "boolean": "bool",
        "timestamp": "string"
    }

    DB_TYPE_TO_MODEL_MAPPING = {
        "integer": "int32",
        "smallint": "int32",
        "string": "sql.NullString",
        "float": "float32",
        "boolean": "bool",
        "timestamp": 'time.Time `sql:"DEFAULT:current_timestamp"`'
    }

    DELIMTER = '$'
    PROTECT_STRING = 'CONTENT_PROTECTED'

    def __init__(self, name, data):
        self.name = name
        self.data = data
        self.skey = 'uid'
        self.pkey = 'id'

        for k,v in self.data.items():
            values = v.split(',')
            if len(values) > 1:
                for each in values:
                    if each == 'pkey':
                        self.pkey = k
                self.data[k] = values[0]

    @classmethod
    def to_snake_format(cls, string):
        s1 = re.sub('(.)([A-Z][a-z]+)', r'\1_\2', string)
        return re.sub('([a-z0-9])([A-Z])', r'\1_\2', s1).lower()

    @classmethod
    def to_camel_format(cls, string):
        camel_format = ''
        if isinstance(string, str):
            for _s_ in string.split('_'):
                camel_format += _s_.capitalize()
        return camel_format

    @classmethod
    def load(cls, name, file):
        file_data = yaml.load(open(file))

        data = {}
        for k,v in file_data.items():
            data[cls.to_snake_format(k)] = v.lower()

        return cls(cls.to_snake_format(name), data)

    def place_holders(self, line):
        p = re.compile(r'\$(\w+)\$')
        result = p.findall(line)

        holders = []
        if result:
            for group in result:
                holders.append(group)
        return holders

    # begin:对应于placehold函数
    def model_name(self):
        return self.name

    def primary_key(self):
        return self.to_camel_format(self.pkey)

    def service_name(self):
        return self.to_camel_format(self.name)

    def table_name(self):
        return self.name

    def get_request_fields(self):
        return '    %s %s = 1;' % (self.DB_TYPE_TO_PROTO_MAPPING[self.data[self.pkey]], self.pkey)

    def get_response_fields(self):
        i = 1
        result = []
        for k, v in self.data.items():
            result.append('    %s %s = %s;'%(self.DB_TYPE_TO_PROTO_MAPPING[v], k, i))
            i += 1
        return '\n'.join(result)

    def get_server_response_fields(self):
        result = []
        for k, v in self.data.items():
            line = '%s: %s.%s'%(self.to_camel_format(k),
                                           self.model_name(),
                                           self.to_camel_format(k))
            if self.DB_TYPE_TO_PROTO_MAPPING[v] == 'string':
                line += '.String,'
            else:
                line += ','

            result.append(line)

        return '\n'.join(result)

    def update_request_fields(self):
        return '    %s %s = 1;' % (self.DB_TYPE_TO_PROTO_MAPPING[self.data[self.pkey]], self.pkey)

    def update_response_fields(self):
        i = 1
        result = []
        for k, v in self.data.items():
            result.append('    %s %s = %s;'%(self.DB_TYPE_TO_PROTO_MAPPING[v], k, i))
            i += 1
        return '\n'.join(result)

    def update_server_response_fields(self):
        result = []
        for k, v in self.data.items():
            line = '%s: %s.%s'%(self.to_camel_format(k),
                                           self.model_name(),
                                           self.to_camel_format(k))
            if self.DB_TYPE_TO_PROTO_MAPPING[v] == 'string':
                line += '.String,'
            else:
                line += ','

            result.append(line)

        return '\n'.join(result)

    def bind_update_fields(self):
        result = []
        for k, v in self.data.items():
            line = '%s: '%self.to_camel_format(k)
            if v == 'string':
                line += 'sql.NullString{update_%s.%s, true},' % (self.model_name(), self.to_camel_format(k))
            else:
                line += 'update_%s.%s,' % (self.model_name(), self.to_camel_format(k))

            result.append(line)

        return '\n'.join(result)

    def bind_create_fields(self):
        result = []
        for k, v in self.data.items():
            line = '%s: '%self.to_camel_format(k)
            if v == 'string':
                line += 'sql.NullString{in.%s, true},' % (self.to_camel_format(k))
            else:
                line += 'in.%s,' % (self.to_camel_format(k))

            result.append(line)

        return '\n'.join(result)

    def model_import(self):
        imports = []
        if 'string' in self.data.values():
            imports.append('"database/sql"')
        if 'timestamp' in self.data.values():
            imports.append('"time"')

        return "\n".join(imports) or ''

    def server_import(self):
        imports = []
        if 'string' in self.data.values():
            imports.append('"database/sql"')

        return "\n".join(imports) or ''

    def model_fields(self):
        result = []
        for k, v in self.data.items():
            line = '%s %s' % (self.to_camel_format(k), self.DB_TYPE_TO_MODEL_MAPPING[v])
            if k == self.pkey:
                line += ' `gorm:"primary_key"`'

            result.append(line)
        return '\n'.join(result)

    # end:对应于placehold函数

    def write(self, src, dest):
        file = open(src)
        try:
            read_file = open(dest, 'r')
            for line in read_file.readlines():
                if self.PROTECT_STRING in line:
                    file.close()
                    read_file.close()
                    return
        except:
            pass

        write_file = open(dest, 'w')
        for line in file.readlines():
            place_holders = self.place_holders(line)
            if place_holders:
                for each in place_holders:
                    func = getattr(self, each.lower())
                    resp = func()
                    line = line.replace('%s%s%s' % (self.DELIMTER, each, self.DELIMTER), resp)
            write_file.writelines(line)
        file.close()
        write_file.close()

    def write_proto(self):
        self.write('template/proto.template', 'proto/%s.proto' % self.name)

    def write_model(self):
        self.write('template/model.template', 'model/%s.go' % self.name)

    def write_server(self):
        self.write('template/server.template', 'server/%s.go' % self.name)

    def write_test(self):
        self.write('template/test.template', 'test/%s.go' % self.name)

def hashopts(pattern):
    try:
        opts, args = getopt.getopt(sys.argv[1:], pattern)
        myopts = {}
        for key, value in opts:
            myopts[key] = value or True
        return myopts
    except Exception:
        usage()


def usage():
    print '''
Description: 自动代码

Usage:
    python %s [options]

OPTIONS:
   -h    help
   -f    file path(default: spec/)
''' % __file__
    sys.exit(1)

SPEC_PATH = 'spec'


def work():
    myopts = hashopts('hif:')
    myopts.get('-h') and usage()
    filename = myopts.get('-f')

    base = SPEC_PATH
    if filename:
        if os.path.isfile(SPEC_PATH + '/' + filename):
            files = [filename]
        else:
            base = '%s/%s' % (base, filename)
            files = os.listdir(base)
    else:
        files = os.listdir(SPEC_PATH)
        files = [each for each in files if 'yaml' in each]


    for each in files:
        full_path = base + '/' + each
        if os.path.isdir(full_path):
            continue

        spec = Spec.load(each.split('.')[0], full_path)
        spec.write_proto()
        spec.write_model()
        spec.write_server()
        spec.write_test()

        print full_path, ' success!'

work()