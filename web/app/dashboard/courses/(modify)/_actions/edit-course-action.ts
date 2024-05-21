"use client"

import { axios } from '@/lib/axios'
import { Session } from 'next-auth'
import React from 'react'

const editCourse = async (session: Session, courseId: string, request: editCourseRequest): Promise<WebResponse | Error> => {
    try {
        const response = await axios.put(`/courses/${courseId}`, request, {
            headers: {
                Authorization: `Bearer ${session.token.access_token}`
            }
        })

        return response.data
    } catch (error: any) {
        return error
    }
}

export default EditCourse
