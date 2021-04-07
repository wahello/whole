// @ts-ignore
/* eslint-disable */
// Code generated by protoc-gen-ts-umi. DO NOT EDIT.
import {request} from 'umi';

const APIService = '/api';
// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.

type Options = {
  [key: string]: any
}

/** Login 常规账户登录接口 /api */
export async function Login(params: UsersV1.CustomerLogin, options?: Options) {
	return request<UsersV1.CustomerLoginRes>(APIService + '/users/customer/login', {
    	method: 'POST',
    	data: {...params},
    	...(options || {}),
	});
}

/** LoginForCode 验证码登录 /api */
export async function LoginForCode(params: UsersV1.CustomerLogin, options?: Options) {
	return request<UsersV1.CustomerLoginRes>(APIService + '/users/customer/login-code', {
    	method: 'POST',
    	data: {...params},
    	...(options || {}),
	});
}

/** Logout 登出 /api */
export async function Logout(params: UsersV1.NullReq, options?: Options) {
	return request<UsersV1.NullReply>(APIService + '/users/customer/logout', {
    	method: 'POST',
    	data: {...params},
    	...(options || {}),
	});
}

/** UserInfo 当前登录用户信息 /api */
export async function UserInfo(params: UsersV1.NullReq, options?: Options) {
	return request<UsersV1.Customer>(APIService + '/users/customer/info', {
    	method: 'GET',
    	params: {...params},
    	...(options || {}),
	});
}

/** Captcha 发送短信/邮箱验证码 /api */
export async function Captcha(params: UsersV1.CaptchaReq, options?: Options) {
	return request<UsersV1.NullReply>(APIService + '/users/customer/captcha', {
    	method: 'GET',
    	params: {...params},
    	...(options || {}),
	});
}

/** List 账户列表 /api */
export async function List(params: UsersV1.CustomerListOption, options?: Options) {
	return request<UsersV1.CustomerList>(APIService + '/users/customer', {
    	method: 'GET',
    	params: {...params},
    	...(options || {}),
	});
}

/** Get 获取账户信息 /api */
export async function Get(params: UsersV1.CustomerGetOption, options?: Options) {
	return request<UsersV1.Customer>(APIService + '/users/customer/{id}', {
    	method: 'GET',
    	params: {...params},
    	...(options || {}),
	});
}

/** Update 更新用户信息 /api */
export async function Update(params: UsersV1.CustomerOption, options?: Options) {
	return request<UsersV1.Customer>(APIService + '/users/customer', {
    	method: 'PUT',
    	data: {...params},
    	...(options || {}),
	});
}

