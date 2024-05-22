"use client"

import { axios } from '@/lib/axios'
import { Session } from 'next-auth'
import React from 'react'

const createCourse = async (session: Session, request: CreateCourseRequest): Promise<WebResponse | Error> => {
    try {
        const response = await axios.post("/courses", request, {
            headers: {
                Authorization: `Bearer ${session.token.access_token}`
            }
        })

        return response.data
    } catch (error: any) {
        return error
    }
}

export default createCourse
