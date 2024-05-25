import "next-auth"
import "next-auth/jwt"

interface Token {
    access_token: string
    access_token_exp: number
    refresh_token: string

    error?: "RefreshAccessTokenError"
}

interface User {
    id?: string | null
    email?: string | null
    name?: string | null
    roles?: Role[] | null
    permissions?: Permission[] | null
}
declare module "next-auth" {
    interface Session {
        token: Token
    }

    interface User extends User {
        token: Token
    }
}

declare module "next-auth/jwt" {
    interface JWT extends Token {
        user: User
    }
}