// @generated by protoc-gen-es v0.1.0 with parameter "target=js"
// @generated from file proto/rtaa/classification/common/v1/predictions.proto (package proto.rtaa.classification.common.v1, syntax proto3)
/* eslint-disable */
/* @ts-nocheck */

import {proto3} from "@bufbuild/protobuf";

/**
 * @generated from message proto.rtaa.classification.common.v1.Prediction
 */
export const Prediction = proto3.makeMessageType(
  "proto.rtaa.classification.common.v1.Prediction",
  () => [
    { no: 1, name: "label", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "score", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 3, name: "prediction", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);
