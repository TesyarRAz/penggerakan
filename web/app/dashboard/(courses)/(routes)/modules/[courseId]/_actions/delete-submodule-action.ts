import { axios } from "@/lib/axios"
import { Session } from "next-auth"

const deleteSubModule = async (session: Session, id: string): Promise<WebResponse> => {
    const response = await axios.delete(`/modules/${id}`, {
        headers: {
            Authorization: `Bearer ${session.token.access_token}`
        }
    })

    return response.data
}

export default deleteSubModule
