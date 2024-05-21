interface WebResponseValidator {
    tag: string
    val: string
}

interface WebResponse {
    errors: Record<String, WebResponseValidator>
    message: string
}

interface PageResponse<T> {
    data : T[]
    paging: {
        prev_cursor: string
        next_cursor: string
    }
}

interface PageRequest {
    cursor: string
    per_page: number
    search: string
}

interface Role {
    id: string
    name: string
}

interface Permission {
    id: string
    name: string
}

interface CourseResponse {
    id: string
    teacher_id: string
    name: string
    image: string
    created_at: string
}

interface CreateCourseRequest {
    teacher_id: string
    name: string
    image: string
}

interface EditCourseRequest {
    name: string
    image: string
}