import { axios } from '@/lib/axios'
import { Session } from 'next-auth'
import React from 'react'

interface ModulePageResponse extends PageResponse<ModuleResponse>{}

const getModules = async (session: Session, courseId: string): Promise<ModulePageResponse> => {
    const response = await axios.get(`/modules/${courseId}`, {
        headers: {
            Authorization: `Bearer ${session.token.access_token}`
        }
    })

    return response.data
}

export default getModules
