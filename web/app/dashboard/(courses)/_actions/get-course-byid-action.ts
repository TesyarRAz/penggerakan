import { axios } from '@/lib/axios'
import { Session } from 'next-auth'
import React from 'react'

const getCourseById = async (session: Session, courseId: string): Promise<CourseResponse> => {
  const response = await axios.get(`/courses/${courseId}`, {
    headers: {
      Authorization: `Bearer ${session.token.access_token}`
    }
  })

  return response.data
}

export default getCourseById
