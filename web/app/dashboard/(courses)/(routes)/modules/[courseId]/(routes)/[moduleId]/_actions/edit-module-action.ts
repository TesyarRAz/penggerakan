import { axios } from '@/lib/axios'
import { Session } from 'next-auth'
import React from 'react'

const editModule = async (session: Session, moduleId: string, request: EditModuleRequest): Promise<ModuleResponse> => {
    const response = await axios.put(`/modules/${moduleId}`, request, {
        headers: {
            Authorization: `Bearer ${session.token.access_token}`
        }
    })

    return response.data
}

export default editModule
