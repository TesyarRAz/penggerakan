import { axios } from '@/lib/axios'
import { Session } from 'next-auth'
import React from 'react'

const deleteTeacher = async (session: Session, teacherId: string): Promise<WebResponse> => {
    const response = await axios.delete(`/teachers/${teacherId}`, {
        headers: {
        Authorization: `Bearer ${session.token.access_token}`
        }
    })
    
    return response.data
}

export default deleteTeacher
