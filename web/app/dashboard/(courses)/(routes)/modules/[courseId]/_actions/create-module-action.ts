import { axios } from '@/lib/axios'
import { Session } from 'next-auth'

const createModule = async (session: Session, courseId: string, request: CreateModuleRequest): Promise<ModuleResponse> => {
    const response = await axios.post(`/modules/${courseId}`, request, {
        headers: {
            Authorization: `Bearer ${session.token.access_token}`
        }
    })

    return response.data
}

export default createModule
