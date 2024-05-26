import { axios } from '@/lib/axios'
import { Session } from 'next-auth'
import React from 'react'

const editSubModule = async (session: Session, moduleId: string, subModuleId: string, request: EditSubModuleRequest): Promise<SubModuleResponse> => {
  const response = await axios.put(`/submodules/${moduleId}/${subModuleId}`, request, {
    headers: {
      Authorization: `Bearer ${session.token.access_token}`
    }
  })

  return response.data
}

export default editSubModule
