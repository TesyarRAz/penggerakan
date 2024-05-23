import { axios } from '@/lib/axios'
import { Session } from 'next-auth'
import React from 'react'

const findModuleById = async (session: Session, courseId: string, moduleId: string): Promise<ModuleResponse> => {
    const response = await axios.get(`/modules/${courseId}/${moduleId}`, {
        headers: {
            Authorization: `Bearer ${session.token.access_token}`
        }
    })

    return response.data
}

export default findModuleById
