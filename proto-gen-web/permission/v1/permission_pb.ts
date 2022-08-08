// @generated by protoc-gen-es v0.0.10 with parameter "target=ts"
// @generated from file permission/v1/permission.proto (package permission.v1, syntax proto3)
/* eslint-disable */
/* @ts-nocheck */

import type {BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage} from "@bufbuild/protobuf";
import {Message, proto3, protoInt64} from "@bufbuild/protobuf";

/**
 * @generated from message permission.v1.CommonUUIDRequest
 */
export class CommonUUIDRequest extends Message<CommonUUIDRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  constructor(data?: PartialMessage<CommonUUIDRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "permission.v1.CommonUUIDRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CommonUUIDRequest {
    return new CommonUUIDRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CommonUUIDRequest {
    return new CommonUUIDRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CommonUUIDRequest {
    return new CommonUUIDRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CommonUUIDRequest | PlainMessage<CommonUUIDRequest> | undefined, b: CommonUUIDRequest | PlainMessage<CommonUUIDRequest> | undefined): boolean {
    return proto3.util.equals(CommonUUIDRequest, a, b);
  }
}

/**
 * @generated from message permission.v1.CommonResponse
 */
export class CommonResponse extends Message<CommonResponse> {
  /**
   * @generated from field: string data = 1;
   */
  data = "";

  constructor(data?: PartialMessage<CommonResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "permission.v1.CommonResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "data", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CommonResponse {
    return new CommonResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CommonResponse {
    return new CommonResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CommonResponse {
    return new CommonResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CommonResponse | PlainMessage<CommonResponse> | undefined, b: CommonResponse | PlainMessage<CommonResponse> | undefined): boolean {
    return proto3.util.equals(CommonResponse, a, b);
  }
}

/**
 * @generated from message permission.v1.FindAllPermissionsRequest
 */
export class FindAllPermissionsRequest extends Message<FindAllPermissionsRequest> {
  /**
   * @generated from field: int64 page = 1;
   */
  page = protoInt64.zero;

  constructor(data?: PartialMessage<FindAllPermissionsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "permission.v1.FindAllPermissionsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "page", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): FindAllPermissionsRequest {
    return new FindAllPermissionsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): FindAllPermissionsRequest {
    return new FindAllPermissionsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): FindAllPermissionsRequest {
    return new FindAllPermissionsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: FindAllPermissionsRequest | PlainMessage<FindAllPermissionsRequest> | undefined, b: FindAllPermissionsRequest | PlainMessage<FindAllPermissionsRequest> | undefined): boolean {
    return proto3.util.equals(FindAllPermissionsRequest, a, b);
  }
}

/**
 * @generated from message permission.v1.FindAllPermissionsResponse
 */
export class FindAllPermissionsResponse extends Message<FindAllPermissionsResponse> {
  /**
   * @generated from field: int64 totalPage = 1;
   */
  totalPage = protoInt64.zero;

  /**
   * @generated from field: int64 currentPage = 2;
   */
  currentPage = protoInt64.zero;

  /**
   * Permissions
   *
   * @generated from field: repeated permission.v1.Permission data = 3;
   */
  data: Permission[] = [];

  constructor(data?: PartialMessage<FindAllPermissionsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "permission.v1.FindAllPermissionsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "totalPage", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "currentPage", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 3, name: "data", kind: "message", T: Permission, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): FindAllPermissionsResponse {
    return new FindAllPermissionsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): FindAllPermissionsResponse {
    return new FindAllPermissionsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): FindAllPermissionsResponse {
    return new FindAllPermissionsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: FindAllPermissionsResponse | PlainMessage<FindAllPermissionsResponse> | undefined, b: FindAllPermissionsResponse | PlainMessage<FindAllPermissionsResponse> | undefined): boolean {
    return proto3.util.equals(FindAllPermissionsResponse, a, b);
  }
}

/**
 * @generated from message permission.v1.Permission
 */
export class Permission extends Message<Permission> {
  /**
   * Output only. Id of the permission.
   *
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * Output only. Name of the permission.
   *
   * @generated from field: string name = 2;
   */
  name = "";

  /**
   * Output only. Email of the permission.
   *
   * @generated from field: string slug = 3;
   */
  slug = "";

  /**
   * Output only. Status of the permission.
   *
   * @generated from field: int32 status = 4;
   */
  status = 0;

  constructor(data?: PartialMessage<Permission>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "permission.v1.Permission";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "slug", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "status", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Permission {
    return new Permission().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Permission {
    return new Permission().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Permission {
    return new Permission().fromJsonString(jsonString, options);
  }

  static equals(a: Permission | PlainMessage<Permission> | undefined, b: Permission | PlainMessage<Permission> | undefined): boolean {
    return proto3.util.equals(Permission, a, b);
  }
}

/**
 * The request create new Permission
 *
 * @generated from message permission.v1.CreatePermissionRequest
 */
export class CreatePermissionRequest extends Message<CreatePermissionRequest> {
  /**
   * @generated from field: string name = 1;
   */
  name = "";

  /**
   * @generated from field: string slug = 2;
   */
  slug = "";

  /**
   * @generated from field: int32 status = 3;
   */
  status = 0;

  constructor(data?: PartialMessage<CreatePermissionRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "permission.v1.CreatePermissionRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "slug", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "status", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreatePermissionRequest {
    return new CreatePermissionRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreatePermissionRequest {
    return new CreatePermissionRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreatePermissionRequest {
    return new CreatePermissionRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreatePermissionRequest | PlainMessage<CreatePermissionRequest> | undefined, b: CreatePermissionRequest | PlainMessage<CreatePermissionRequest> | undefined): boolean {
    return proto3.util.equals(CreatePermissionRequest, a, b);
  }
}

/**
 * The request update Permission
 *
 * @generated from message permission.v1.UpdatePermissionRequest
 */
export class UpdatePermissionRequest extends Message<UpdatePermissionRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: optional string name = 2;
   */
  name?: string;

  /**
   * @generated from field: optional string slug = 3;
   */
  slug?: string;

  /**
   * @generated from field: optional int32 status = 4;
   */
  status?: number;

  constructor(data?: PartialMessage<UpdatePermissionRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "permission.v1.UpdatePermissionRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 3, name: "slug", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 4, name: "status", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdatePermissionRequest {
    return new UpdatePermissionRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdatePermissionRequest {
    return new UpdatePermissionRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdatePermissionRequest {
    return new UpdatePermissionRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdatePermissionRequest | PlainMessage<UpdatePermissionRequest> | undefined, b: UpdatePermissionRequest | PlainMessage<UpdatePermissionRequest> | undefined): boolean {
    return proto3.util.equals(UpdatePermissionRequest, a, b);
  }
}
