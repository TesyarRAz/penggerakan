import { axios } from "@/lib/axios"
import { Session } from "next-auth"

const createSubModule = async (session: Session, moduleId: string, request: CreateSubModuleRequest): Promise<SubModuleResponse> => {
    const response = await axios.post(`/submodules/${moduleId}`, request, {
        headers: {
            Authorization: `Bearer ${session.token.access_token}`
        }
    })

    return response.data
}

export default createSubModule