# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: circuitbreaker.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x14\x63ircuitbreaker.proto\x12\x08protobuf\x1a\x1bgoogle/protobuf/empty.proto\"\x95\x02\n\x0cRequestInput\x12\x0e\n\x06method\x18\x01 \x01(\t\x12\x0b\n\x03url\x18\x02 \x01(\t\x12\x32\n\x06header\x18\x03 \x03(\x0b\x32\".protobuf.RequestInput.HeaderEntry\x12\x0c\n\x04\x62ody\x18\x04 \x01(\x0c\x12\x41\n\x0einitial_header\x18\x05 \x03(\x0b\x32).protobuf.RequestInput.InitialHeaderEntry\x1a-\n\x0bHeaderEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x1a\x34\n\x12InitialHeaderEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"\xeb\x01\n\x08GetInput\x12\x0b\n\x03url\x18\x01 \x01(\t\x12.\n\x06header\x18\x02 \x03(\x0b\x32\x1e.protobuf.GetInput.HeaderEntry\x12=\n\x0einitial_header\x18\x03 \x03(\x0b\x32%.protobuf.GetInput.InitialHeaderEntry\x1a-\n\x0bHeaderEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x1a\x34\n\x12InitialHeaderEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"\xfc\x01\n\tPostInput\x12\x0b\n\x03url\x18\x01 \x01(\t\x12/\n\x06header\x18\x02 \x03(\x0b\x32\x1f.protobuf.PostInput.HeaderEntry\x12\x0c\n\x04\x62ody\x18\x03 \x01(\x0c\x12>\n\x0einitial_header\x18\x04 \x03(\x0b\x32&.protobuf.PostInput.InitialHeaderEntry\x1a-\n\x0bHeaderEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x1a\x34\n\x12InitialHeaderEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"\xf9\x01\n\x08PutInput\x12\x0b\n\x03url\x18\x01 \x01(\t\x12.\n\x06header\x18\x02 \x03(\x0b\x32\x1e.protobuf.PutInput.HeaderEntry\x12\x0c\n\x04\x62ody\x18\x03 \x01(\x0c\x12=\n\x0einitial_header\x18\x04 \x03(\x0b\x32%.protobuf.PutInput.InitialHeaderEntry\x1a-\n\x0bHeaderEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x1a\x34\n\x12InitialHeaderEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"\xf4\x01\n\x0b\x44\x65leteInput\x12\x0b\n\x03url\x18\x01 \x01(\t\x12\x31\n\x06header\x18\x02 \x03(\x0b\x32!.protobuf.DeleteInput.HeaderEntry\x12@\n\x0einitial_header\x18\x03 \x03(\x0b\x32(.protobuf.DeleteInput.InitialHeaderEntry\x1a-\n\x0bHeaderEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x1a\x34\n\x12InitialHeaderEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"\xed\x01\n\x08Response\x12\x0e\n\x06status\x18\x01 \x01(\t\x12\x13\n\x0bstatus_code\x18\x02 \x01(\x05\x12\r\n\x05proto\x18\x03 \x01(\t\x12\x13\n\x0bproto_major\x18\x04 \x01(\x05\x12\x13\n\x0bproto_minor\x18\x05 \x01(\x05\x12.\n\x06header\x18\x06 \x03(\x0b\x32\x1e.protobuf.Response.HeaderEntry\x12\x0c\n\x04\x62ody\x18\x07 \x01(\x0c\x12\x16\n\x0e\x63ontent_length\x18\x08 \x01(\x03\x1a-\n\x0bHeaderEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"G\n\x10ServiceEndpoints\x12\x10\n\x08services\x18\x01 \x03(\t\x12\x11\n\tendpoints\x18\x02 \x03(\t\x12\x0e\n\x06\x65xpiry\x18\x03 \x01(\x03\x32\x8b\x02\n\x04Http\x12\x37\n\x07Request\x12\x16.protobuf.RequestInput\x1a\x12.protobuf.Response\"\x00\x12/\n\x03Get\x12\x12.protobuf.GetInput\x1a\x12.protobuf.Response\"\x00\x12\x31\n\x04Post\x12\x13.protobuf.PostInput\x1a\x12.protobuf.Response\"\x00\x12/\n\x03Put\x12\x12.protobuf.PutInput\x1a\x12.protobuf.Response\"\x00\x12\x35\n\x06\x44\x65lete\x12\x15.protobuf.DeleteInput\x1a\x12.protobuf.Response\"\x00\x32\x97\x01\n\x08Listener\x12\x44\n\x0cOpenCircuits\x12\x1a.protobuf.ServiceEndpoints\x1a\x16.google.protobuf.Empty\"\x00\x12\x45\n\rCloseCircuits\x12\x1a.protobuf.ServiceEndpoints\x1a\x16.google.protobuf.Empty\"\x00\x42%Z#github.com/naufal-dean/ccb/protobufb\x06proto3')



_REQUESTINPUT = DESCRIPTOR.message_types_by_name['RequestInput']
_REQUESTINPUT_HEADERENTRY = _REQUESTINPUT.nested_types_by_name['HeaderEntry']
_REQUESTINPUT_INITIALHEADERENTRY = _REQUESTINPUT.nested_types_by_name['InitialHeaderEntry']
_GETINPUT = DESCRIPTOR.message_types_by_name['GetInput']
_GETINPUT_HEADERENTRY = _GETINPUT.nested_types_by_name['HeaderEntry']
_GETINPUT_INITIALHEADERENTRY = _GETINPUT.nested_types_by_name['InitialHeaderEntry']
_POSTINPUT = DESCRIPTOR.message_types_by_name['PostInput']
_POSTINPUT_HEADERENTRY = _POSTINPUT.nested_types_by_name['HeaderEntry']
_POSTINPUT_INITIALHEADERENTRY = _POSTINPUT.nested_types_by_name['InitialHeaderEntry']
_PUTINPUT = DESCRIPTOR.message_types_by_name['PutInput']
_PUTINPUT_HEADERENTRY = _PUTINPUT.nested_types_by_name['HeaderEntry']
_PUTINPUT_INITIALHEADERENTRY = _PUTINPUT.nested_types_by_name['InitialHeaderEntry']
_DELETEINPUT = DESCRIPTOR.message_types_by_name['DeleteInput']
_DELETEINPUT_HEADERENTRY = _DELETEINPUT.nested_types_by_name['HeaderEntry']
_DELETEINPUT_INITIALHEADERENTRY = _DELETEINPUT.nested_types_by_name['InitialHeaderEntry']
_RESPONSE = DESCRIPTOR.message_types_by_name['Response']
_RESPONSE_HEADERENTRY = _RESPONSE.nested_types_by_name['HeaderEntry']
_SERVICEENDPOINTS = DESCRIPTOR.message_types_by_name['ServiceEndpoints']
RequestInput = _reflection.GeneratedProtocolMessageType('RequestInput', (_message.Message,), {

  'HeaderEntry' : _reflection.GeneratedProtocolMessageType('HeaderEntry', (_message.Message,), {
    'DESCRIPTOR' : _REQUESTINPUT_HEADERENTRY,
    '__module__' : 'circuitbreaker_pb2'
    # @@protoc_insertion_point(class_scope:protobuf.RequestInput.HeaderEntry)
    })
  ,

  'InitialHeaderEntry' : _reflection.GeneratedProtocolMessageType('InitialHeaderEntry', (_message.Message,), {
    'DESCRIPTOR' : _REQUESTINPUT_INITIALHEADERENTRY,
    '__module__' : 'circuitbreaker_pb2'
    # @@protoc_insertion_point(class_scope:protobuf.RequestInput.InitialHeaderEntry)
    })
  ,
  'DESCRIPTOR' : _REQUESTINPUT,
  '__module__' : 'circuitbreaker_pb2'
  # @@protoc_insertion_point(class_scope:protobuf.RequestInput)
  })
_sym_db.RegisterMessage(RequestInput)
_sym_db.RegisterMessage(RequestInput.HeaderEntry)
_sym_db.RegisterMessage(RequestInput.InitialHeaderEntry)

GetInput = _reflection.GeneratedProtocolMessageType('GetInput', (_message.Message,), {

  'HeaderEntry' : _reflection.GeneratedProtocolMessageType('HeaderEntry', (_message.Message,), {
    'DESCRIPTOR' : _GETINPUT_HEADERENTRY,
    '__module__' : 'circuitbreaker_pb2'
    # @@protoc_insertion_point(class_scope:protobuf.GetInput.HeaderEntry)
    })
  ,

  'InitialHeaderEntry' : _reflection.GeneratedProtocolMessageType('InitialHeaderEntry', (_message.Message,), {
    'DESCRIPTOR' : _GETINPUT_INITIALHEADERENTRY,
    '__module__' : 'circuitbreaker_pb2'
    # @@protoc_insertion_point(class_scope:protobuf.GetInput.InitialHeaderEntry)
    })
  ,
  'DESCRIPTOR' : _GETINPUT,
  '__module__' : 'circuitbreaker_pb2'
  # @@protoc_insertion_point(class_scope:protobuf.GetInput)
  })
_sym_db.RegisterMessage(GetInput)
_sym_db.RegisterMessage(GetInput.HeaderEntry)
_sym_db.RegisterMessage(GetInput.InitialHeaderEntry)

PostInput = _reflection.GeneratedProtocolMessageType('PostInput', (_message.Message,), {

  'HeaderEntry' : _reflection.GeneratedProtocolMessageType('HeaderEntry', (_message.Message,), {
    'DESCRIPTOR' : _POSTINPUT_HEADERENTRY,
    '__module__' : 'circuitbreaker_pb2'
    # @@protoc_insertion_point(class_scope:protobuf.PostInput.HeaderEntry)
    })
  ,

  'InitialHeaderEntry' : _reflection.GeneratedProtocolMessageType('InitialHeaderEntry', (_message.Message,), {
    'DESCRIPTOR' : _POSTINPUT_INITIALHEADERENTRY,
    '__module__' : 'circuitbreaker_pb2'
    # @@protoc_insertion_point(class_scope:protobuf.PostInput.InitialHeaderEntry)
    })
  ,
  'DESCRIPTOR' : _POSTINPUT,
  '__module__' : 'circuitbreaker_pb2'
  # @@protoc_insertion_point(class_scope:protobuf.PostInput)
  })
_sym_db.RegisterMessage(PostInput)
_sym_db.RegisterMessage(PostInput.HeaderEntry)
_sym_db.RegisterMessage(PostInput.InitialHeaderEntry)

PutInput = _reflection.GeneratedProtocolMessageType('PutInput', (_message.Message,), {

  'HeaderEntry' : _reflection.GeneratedProtocolMessageType('HeaderEntry', (_message.Message,), {
    'DESCRIPTOR' : _PUTINPUT_HEADERENTRY,
    '__module__' : 'circuitbreaker_pb2'
    # @@protoc_insertion_point(class_scope:protobuf.PutInput.HeaderEntry)
    })
  ,

  'InitialHeaderEntry' : _reflection.GeneratedProtocolMessageType('InitialHeaderEntry', (_message.Message,), {
    'DESCRIPTOR' : _PUTINPUT_INITIALHEADERENTRY,
    '__module__' : 'circuitbreaker_pb2'
    # @@protoc_insertion_point(class_scope:protobuf.PutInput.InitialHeaderEntry)
    })
  ,
  'DESCRIPTOR' : _PUTINPUT,
  '__module__' : 'circuitbreaker_pb2'
  # @@protoc_insertion_point(class_scope:protobuf.PutInput)
  })
_sym_db.RegisterMessage(PutInput)
_sym_db.RegisterMessage(PutInput.HeaderEntry)
_sym_db.RegisterMessage(PutInput.InitialHeaderEntry)

DeleteInput = _reflection.GeneratedProtocolMessageType('DeleteInput', (_message.Message,), {

  'HeaderEntry' : _reflection.GeneratedProtocolMessageType('HeaderEntry', (_message.Message,), {
    'DESCRIPTOR' : _DELETEINPUT_HEADERENTRY,
    '__module__' : 'circuitbreaker_pb2'
    # @@protoc_insertion_point(class_scope:protobuf.DeleteInput.HeaderEntry)
    })
  ,

  'InitialHeaderEntry' : _reflection.GeneratedProtocolMessageType('InitialHeaderEntry', (_message.Message,), {
    'DESCRIPTOR' : _DELETEINPUT_INITIALHEADERENTRY,
    '__module__' : 'circuitbreaker_pb2'
    # @@protoc_insertion_point(class_scope:protobuf.DeleteInput.InitialHeaderEntry)
    })
  ,
  'DESCRIPTOR' : _DELETEINPUT,
  '__module__' : 'circuitbreaker_pb2'
  # @@protoc_insertion_point(class_scope:protobuf.DeleteInput)
  })
_sym_db.RegisterMessage(DeleteInput)
_sym_db.RegisterMessage(DeleteInput.HeaderEntry)
_sym_db.RegisterMessage(DeleteInput.InitialHeaderEntry)

Response = _reflection.GeneratedProtocolMessageType('Response', (_message.Message,), {

  'HeaderEntry' : _reflection.GeneratedProtocolMessageType('HeaderEntry', (_message.Message,), {
    'DESCRIPTOR' : _RESPONSE_HEADERENTRY,
    '__module__' : 'circuitbreaker_pb2'
    # @@protoc_insertion_point(class_scope:protobuf.Response.HeaderEntry)
    })
  ,
  'DESCRIPTOR' : _RESPONSE,
  '__module__' : 'circuitbreaker_pb2'
  # @@protoc_insertion_point(class_scope:protobuf.Response)
  })
_sym_db.RegisterMessage(Response)
_sym_db.RegisterMessage(Response.HeaderEntry)

ServiceEndpoints = _reflection.GeneratedProtocolMessageType('ServiceEndpoints', (_message.Message,), {
  'DESCRIPTOR' : _SERVICEENDPOINTS,
  '__module__' : 'circuitbreaker_pb2'
  # @@protoc_insertion_point(class_scope:protobuf.ServiceEndpoints)
  })
_sym_db.RegisterMessage(ServiceEndpoints)

_HTTP = DESCRIPTOR.services_by_name['Http']
_LISTENER = DESCRIPTOR.services_by_name['Listener']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z#github.com/naufal-dean/ccb/protobuf'
  _REQUESTINPUT_HEADERENTRY._options = None
  _REQUESTINPUT_HEADERENTRY._serialized_options = b'8\001'
  _REQUESTINPUT_INITIALHEADERENTRY._options = None
  _REQUESTINPUT_INITIALHEADERENTRY._serialized_options = b'8\001'
  _GETINPUT_HEADERENTRY._options = None
  _GETINPUT_HEADERENTRY._serialized_options = b'8\001'
  _GETINPUT_INITIALHEADERENTRY._options = None
  _GETINPUT_INITIALHEADERENTRY._serialized_options = b'8\001'
  _POSTINPUT_HEADERENTRY._options = None
  _POSTINPUT_HEADERENTRY._serialized_options = b'8\001'
  _POSTINPUT_INITIALHEADERENTRY._options = None
  _POSTINPUT_INITIALHEADERENTRY._serialized_options = b'8\001'
  _PUTINPUT_HEADERENTRY._options = None
  _PUTINPUT_HEADERENTRY._serialized_options = b'8\001'
  _PUTINPUT_INITIALHEADERENTRY._options = None
  _PUTINPUT_INITIALHEADERENTRY._serialized_options = b'8\001'
  _DELETEINPUT_HEADERENTRY._options = None
  _DELETEINPUT_HEADERENTRY._serialized_options = b'8\001'
  _DELETEINPUT_INITIALHEADERENTRY._options = None
  _DELETEINPUT_INITIALHEADERENTRY._serialized_options = b'8\001'
  _RESPONSE_HEADERENTRY._options = None
  _RESPONSE_HEADERENTRY._serialized_options = b'8\001'
  _REQUESTINPUT._serialized_start=64
  _REQUESTINPUT._serialized_end=341
  _REQUESTINPUT_HEADERENTRY._serialized_start=242
  _REQUESTINPUT_HEADERENTRY._serialized_end=287
  _REQUESTINPUT_INITIALHEADERENTRY._serialized_start=289
  _REQUESTINPUT_INITIALHEADERENTRY._serialized_end=341
  _GETINPUT._serialized_start=344
  _GETINPUT._serialized_end=579
  _GETINPUT_HEADERENTRY._serialized_start=242
  _GETINPUT_HEADERENTRY._serialized_end=287
  _GETINPUT_INITIALHEADERENTRY._serialized_start=289
  _GETINPUT_INITIALHEADERENTRY._serialized_end=341
  _POSTINPUT._serialized_start=582
  _POSTINPUT._serialized_end=834
  _POSTINPUT_HEADERENTRY._serialized_start=242
  _POSTINPUT_HEADERENTRY._serialized_end=287
  _POSTINPUT_INITIALHEADERENTRY._serialized_start=289
  _POSTINPUT_INITIALHEADERENTRY._serialized_end=341
  _PUTINPUT._serialized_start=837
  _PUTINPUT._serialized_end=1086
  _PUTINPUT_HEADERENTRY._serialized_start=242
  _PUTINPUT_HEADERENTRY._serialized_end=287
  _PUTINPUT_INITIALHEADERENTRY._serialized_start=289
  _PUTINPUT_INITIALHEADERENTRY._serialized_end=341
  _DELETEINPUT._serialized_start=1089
  _DELETEINPUT._serialized_end=1333
  _DELETEINPUT_HEADERENTRY._serialized_start=242
  _DELETEINPUT_HEADERENTRY._serialized_end=287
  _DELETEINPUT_INITIALHEADERENTRY._serialized_start=289
  _DELETEINPUT_INITIALHEADERENTRY._serialized_end=341
  _RESPONSE._serialized_start=1336
  _RESPONSE._serialized_end=1573
  _RESPONSE_HEADERENTRY._serialized_start=242
  _RESPONSE_HEADERENTRY._serialized_end=287
  _SERVICEENDPOINTS._serialized_start=1575
  _SERVICEENDPOINTS._serialized_end=1646
  _HTTP._serialized_start=1649
  _HTTP._serialized_end=1916
  _LISTENER._serialized_start=1919
  _LISTENER._serialized_end=2070
# @@protoc_insertion_point(module_scope)