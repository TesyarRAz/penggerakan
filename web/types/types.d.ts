interface WebResponseValidator {
    tag: string
    val: string
}

interface WebResponse {
    errors: Record<String, WebResponseValidator>
    message: string
}

interface PageResponse<T> {
    data: T[]
    paging: {
        prev_cursor: string
        next_cursor: string
    }
}

interface PageRequest {
    cursor?: string
    per_page?: number
    search?: string
}

interface Role {
    id: string
    name: string
}

interface Permission {
    id: string
    name: string
}

interface LoginResponse {
    user: User
    token: Token
}

interface RefreshTokenResponse {
    access_token: string
    access_token_exp: number
    refresh_token: string
}

interface CourseResponse {
    id: string
    teacher_id: string
    name: string
    image: string
    created_at?: string
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

interface UserResponse {
    id: string
    name: string
    profile_image: string
    email: string
    roles: Role[]
    permissions: Permission[]
    created_at?: string
}

interface ModuleResponse {
    id: string
    course_id: string
    name: string
    created_at?: string
}

interface CreateModuleRequest {
    name: string
}

interface EditModuleRequest {
    name: string
}

interface SubModuleResponse {
    id: string
    module_id: string
    name: string
    structure: Record<string, any>
    created_at?: string
}

interface TeacherResponse {
    id: string
    user_id: string
    name: string
    email: string
    profile_image: string
    created_at?: string
}

interface CreateTeacherRequest {
    user_id: string
    name: string
    profile_image: string
}

interface EditTeacherRequest {
    name: string
    profile_image: string
}

interface SubModuleStructure {
    resource_id?: string
    resource_type?: string
    label?: string
    value?: string
    children?: SubModuleStructure[]
}
