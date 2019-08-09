# This file is autogenerated. Do not edit it by hand. Regenerate it with:
#   srb rbi gems

# typed: ignore
#
# If you would like to make changes to this file, great! Please create the gem's shim here:
#
#   https://github.com/sorbet/sorbet-typed/new/master?filename=lib/google-protobuf-3.7.1-universal/all/google-protobuf-3.7.1-universal.rbi
#
# google-protobuf-3.7.1-universal-darwin
module Google
end
module Google::Protobuf
  def self.decode(klass, proto); end
  def self.decode_json(klass, json, options = nil); end
  def self.deep_copy(arg0); end
  def self.discard_unknown(arg0); end
  def self.encode(msg); end
  def self.encode_json(msg, options = nil); end
end
module Google::Protobuf::MessageExts
  def self.included(klass); end
  def to_json(options = nil); end
  def to_proto; end
end
module Google::Protobuf::MessageExts::ClassMethods
end
class Google::Protobuf::DescriptorPool
  def add(arg0); end
  def build(*arg0); end
  def lookup(arg0); end
  def self.generated_pool; end
end
class Google::Protobuf::Descriptor
  def add_field(arg0); end
  def add_oneof(arg0); end
  def each; end
  def each_oneof; end
  def file_descriptor; end
  def initialize(arg0); end
  def lookup(arg0); end
  def lookup_oneof(arg0); end
  def msgclass; end
  def name; end
  def name=(arg0); end
end
class Google::Protobuf::FileDescriptor
  def initialize(*arg0); end
  def name; end
  def syntax; end
  def syntax=(arg0); end
end
class Google::Protobuf::FieldDescriptor
  def clear(arg0); end
  def default; end
  def default=(arg0); end
  def get(arg0); end
  def has?(arg0); end
  def label; end
  def label=(arg0); end
  def name; end
  def name=(arg0); end
  def number; end
  def number=(arg0); end
  def set(arg0, arg1); end
  def submsg_name; end
  def submsg_name=(arg0); end
  def subtype; end
  def type; end
  def type=(arg0); end
end
class Google::Protobuf::OneofDescriptor
  def add_field(arg0); end
  def each; end
  def name; end
  def name=(arg0); end
end
class Google::Protobuf::EnumDescriptor
  def add_value(arg0, arg1); end
  def each; end
  def enummodule; end
  def file_descriptor; end
  def initialize(arg0); end
  def lookup_name(arg0); end
  def lookup_value(arg0); end
  def name; end
  def name=(arg0); end
end
class Google::Protobuf::Internal::MessageBuilderContext
  def initialize(arg0, arg1); end
  def map(*arg0); end
  def oneof(arg0); end
  def optional(*arg0); end
  def repeated(*arg0); end
  def required(*arg0); end
end
class Google::Protobuf::Internal::OneofBuilderContext
  def initialize(arg0, arg1); end
  def optional(*arg0); end
end
class Google::Protobuf::Internal::EnumBuilderContext
  def initialize(arg0); end
  def value(arg0, arg1); end
end
class Google::Protobuf::Internal::FileBuilderContext
  def add_enum(arg0); end
  def add_message(arg0); end
  def initialize(arg0, arg1); end
end
class Google::Protobuf::Internal::Builder
  def add_enum(arg0); end
  def add_file(*arg0); end
  def add_message(arg0); end
  def finalize_to_pool(arg0); end
  def initialize; end
end
class Google::Protobuf::RepeatedField
  def &(*args, &block); end
  def *(*args, &block); end
  def +(arg0); end
  def -(*args, &block); end
  def <<(arg0); end
  def <=>(*args, &block); end
  def ==(arg0); end
  def [](*arg0); end
  def []=(arg0, arg1); end
  def assoc(*args, &block); end
  def at(*arg0); end
  def bsearch(*args, &block); end
  def bsearch_index(*args, &block); end
  def clear; end
  def clone; end
  def collect!(*args, &block); end
  def combination(*args, &block); end
  def compact!(*args, &block); end
  def compact(*args, &block); end
  def concat(arg0); end
  def count(*args, &block); end
  def cycle(*args, &block); end
  def delete(*args, &block); end
  def delete_at(*args, &block); end
  def delete_if(*args, &block); end
  def dig(*args, &block); end
  def drop(*args, &block); end
  def drop_while(*args, &block); end
  def dup; end
  def each; end
  def each_index(*arg0); end
  def empty?; end
  def eql?(*args, &block); end
  def fetch(*args, &block); end
  def fill(*args, &block); end
  def find_index(*args, &block); end
  def first(n = nil); end
  def flatten!(*args, &block); end
  def flatten(*args, &block); end
  def hash; end
  def include?(*args, &block); end
  def index(*args, &block); end
  def initialize(*arg0); end
  def insert(*args, &block); end
  def inspect(*args, &block); end
  def join(*args, &block); end
  def keep_if(*args, &block); end
  def last(n = nil); end
  def length; end
  def map!(*args, &block); end
  def map; end
  def pack(*args, &block); end
  def permutation(*args, &block); end
  def pop(n = nil); end
  def pop_one; end
  def pretty_print(*args, &block); end
  def pretty_print_cycle(*args, &block); end
  def product(*args, &block); end
  def push(arg0); end
  def rassoc(*args, &block); end
  def reject!(*args, &block); end
  def repeated_combination(*args, &block); end
  def repeated_permutation(*args, &block); end
  def replace(arg0); end
  def reverse!(*args, &block); end
  def reverse(*args, &block); end
  def rindex(*args, &block); end
  def rotate!(*args, &block); end
  def rotate(*args, &block); end
  def sample(*args, &block); end
  def select!(*args, &block); end
  def self.define_array_wrapper_method(method_name); end
  def self.define_array_wrapper_with_result_method(method_name); end
  def shelljoin(*args, &block); end
  def shift(*args, &block); end
  def shuffle!(*args, &block); end
  def shuffle(*args, &block); end
  def size; end
  def slice!(*args, &block); end
  def slice(*arg0); end
  def sort!(*args, &block); end
  def sort_by!(*args, &block); end
  def to_ary; end
  def to_s(*args, &block); end
  def transpose(*args, &block); end
  def uniq!(*args, &block); end
  def uniq(*args, &block); end
  def unshift(*args, &block); end
  def values_at; end
  def |(*args, &block); end
  extend Forwardable
end
class Google::Protobuf::Map
  def ==(arg0); end
  def [](arg0); end
  def []=(arg0, arg1); end
  def clear; end
  def delete(arg0); end
  def dup; end
  def each; end
  def has_key?(arg0); end
  def hash; end
  def initialize(*arg0); end
  def inspect; end
  def keys; end
  def length; end
  def merge(arg0); end
  def to_h; end
  def to_hash; end
  def values; end
end
class Anonymous_Struct_1 < Struct
  def external_enumerator; end
  def external_enumerator=(_); end
  def repeated_field; end
  def repeated_field=(_); end
  def self.[](*arg0); end
  def self.inspect; end
  def self.members; end
  def self.new(*arg0); end
end
class Google::Protobuf::RepeatedField::ProxyingEnumerator < Anonymous_Struct_1
  def each(*args, &block); end
end
class Google::Protobuf::Error < StandardError
end
class Google::Protobuf::ParseError < Google::Protobuf::Error
end
class Google::Protobuf::TypeError < TypeError
end