import { axios } from '@/lib/axios'
import { Session } from 'next-auth'
import React from 'react'

const deleteUser = async (session: Session, userId: string): Promise<WebResponse | Error> => {
    try {
        const response = await axios.delete(`/users/${userId}`, {
            headers: {
                Authorization: `Bearer ${session.token.access_token}`
            }
        })
    
        return response.data
    } catch (error: any) {
        return error
    }
}

export default deleteUser
