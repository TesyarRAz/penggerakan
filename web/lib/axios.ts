import Axios, { CreateAxiosDefaults } from 'axios';

export const AxiosAppConfig: CreateAxiosDefaults = {
  baseURL: process.env.API_URL,
  headers: {
    'Content-Type': 'application/json',
    'Accept': 'application/json',
  },
}

export const axios = Axios.create(AxiosAppConfig);
export const axiosAuth = Axios.create(AxiosAppConfig);