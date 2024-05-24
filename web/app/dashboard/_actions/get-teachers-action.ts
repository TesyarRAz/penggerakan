import { axios } from '@/lib/axios'
import { Session } from 'next-auth'
import React from 'react'

interface TeacherPageResponse extends PageResponse<TeacherResponse> { }

const getTeachers = async (session: Session, request?: PageRequest): Promise<TeacherPageResponse> => {
  const response = await axios.get("/teachers", {
    params: request,
    headers: {
      Authorization: `Bearer ${session.token.access_token}`
    }
  })

  return response.data
}



export default getTeachers
