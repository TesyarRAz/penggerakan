import { axios } from '@/lib/axios'
import { Session } from 'next-auth'
import React from 'react'

interface UserPageResponse extends PageResponse<UserResponse> { }

const getUsers = async (session: Session, request?: PageRequest): Promise<UserPageResponse> => {
  const response = await axios.get("/users", {
    params: request,
    headers: {
      Authorization: `Bearer ${session.token.access_token}`
    }
  })

  return response.data
}

export default getUsers
