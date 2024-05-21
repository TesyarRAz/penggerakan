"use client"

import { axios } from '@/lib/axios'
import { Session } from 'next-auth'
import React from 'react'

const deleteCourseAction = async (session: Session, courseId: string): Promise<WebResponse | Error> => {
  try {
    const response = await axios.delete(`/courses/${courseId}`, {
        headers: {
            Authorization: `Bearer ${session.token.access_token}`
        }
    })

    return response.data
  } catch (error: any) {
    return error
  }
}

export default deleteCourseAction
