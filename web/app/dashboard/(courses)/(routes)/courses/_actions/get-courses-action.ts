import { axios } from "@/lib/axios"
import { Session } from "next-auth"

interface CoursePageResponse extends PageResponse<CourseResponse>{}

export const getCourses = async (session: Session, request?: PageRequest): Promise<CoursePageResponse> => {
    const response = await axios.get("/courses", {
        params: request,
        headers: {
            Authorization: `Bearer ${session.token.access_token}`
        }
    })
    
    return response.data
}