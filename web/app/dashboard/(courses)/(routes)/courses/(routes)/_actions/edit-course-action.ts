"use client"

import { axios } from '@/lib/axios'
import { Session } from 'next-auth'
import React from 'react'

const editCourse = async (session: Session, courseId: string, request: EditCourseRequest): Promise<WebResponse> => {
    const response = await axios.put(`/courses/${courseId}`, request, {
        headers: {
            Authorization: `Bearer ${session.token.access_token}`
        }
    })

    return response.data
}

export default editCourse
