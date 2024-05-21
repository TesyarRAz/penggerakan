"use client"

import axios from 'axios'
import { signIn, useSession } from 'next-auth/react'
import React, { useEffect, useRef } from 'react'
import { axiosAuth } from '../axios'

const useAxiosAuth = () => {
  const { data: session } = useSession()
  
  useEffect(() => {
    const requestInterceptor = axiosAuth.interceptors.request.use(
      (config) => {
        if (!config.headers["Authorization"]) {
          config.headers["Authorization"] = `Bearer ${session?.token.access_token}`
        }
        return config
      },
      (error) => Promise.reject(error)
    )

    const responseInterceptor = axiosAuth.interceptors.response.use(
      (response) => response,
      async (error) => {
        const previousRequest = error?.config
        if (error.response.status === 401 && !previousRequest?.sent) {
          previousRequest.sent = true
          await signIn("app-credentials", { redirect: false })
          return axiosAuth(previousRequest)
        }
        return Promise.reject(error)
      }
    )

    return () => {
      axiosAuth.interceptors.request.eject(requestInterceptor)
      axiosAuth.interceptors.response.eject(responseInterceptor)
    }
  }, [session])
}

export default useAxiosAuth
