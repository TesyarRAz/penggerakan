import { axios } from '@/lib/axios'
import { Session } from 'next-auth'
import React from 'react'

interface SubModulesPageResponse extends PageResponse<SubModuleResponse>{}

const getSubModules = async (session: Session, moduleId: string): Promise<SubModulesPageResponse> => {
  const response = await axios.get(`/submodules/${moduleId}`, {
    headers: {
        Authorization: `Bearer ${session.token.access_token}`
    }
  })

    return response.data
}

export default getSubModules
