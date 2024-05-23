import { axios } from '@/lib/axios'
import { Session } from 'next-auth'
import React from 'react'

const deleteModule = async (session: Session, id: string): Promise<WebResponse> => {
    const response = await axios.delete(`/modules/${id}`, {
        headers: {
            Authorization: `Bearer ${session.token.access_token}`
        }
    })

    return response.data
}

export default deleteModule
