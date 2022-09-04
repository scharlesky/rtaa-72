// @generated by protoc-gen-connect-web v0.1.0 with parameter "target=js"
// @generated from file proto/rtaa/classification/comments/v1/comments.proto (package proto.rtaa.classification.comments.v1, syntax proto3)
/* eslint-disable */
/* @ts-nocheck */

import {Comment, CommentList, ResponseToxic, ResponseToxicList} from "./comments_pb.js";
import {MethodKind} from "@bufbuild/protobuf";

/**
 * @generated from service proto.rtaa.classification.comments.v1.CommentService
 */
export const CommentService = {
  typeName: "proto.rtaa.classification.comments.v1.CommentService",
  methods: {
    /**
     * @generated from rpc proto.rtaa.classification.comments.v1.CommentService.ClassifyToxic
     */
    classifyToxic: {
      name: "ClassifyToxic",
      I: Comment,
      O: ResponseToxic,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.rtaa.classification.comments.v1.CommentService.ClassifyToxicList
     */
    classifyToxicList: {
      name: "ClassifyToxicList",
      I: CommentList,
      O: ResponseToxicList,
      kind: MethodKind.Unary,
    },
  }
};

