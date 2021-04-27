// @ts-ignore
/* eslint-disable */
// Code generated by protoc-gen-ts-umi. DO NOT EDIT.

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.

declare namespace NewsV1 {
	/** MessageList */
	type MessageList = {
		list?:Array<NewsV1.Message>
		total?:number
		page?:number
		page_size?:number
	}
	/** MessageListOption */
	type MessageListOption = {
		type?:number
		read?:number
		total?:number
		unread_total?:number
		page?:number
		page_size?:number
	}
	/** Message */
	type Message = {
		type?:string
		title?:string
		content?:string
		to?:string
		read?:boolean
		birthday?:GoogleProtobuf.Timestamp
	}
	/** MessageIds */
	type MessageIds = {
		id?:number
	}
	/** MessageCreateOption */
	type MessageCreateOption = {
		type?:string
		title?:string
		content?:string
		to?:string
	}
	/** MessageDeleteOption */
	type MessageDeleteOption = {
		id?:number
	}
}

declare namespace GoogleProtobuf {
	/** Timestamp */
	type Timestamp = {
		seconds?:number
		nanos?:number
	}
}

