import "next-auth"
import "next-auth/jwt"

interface Token {
    access_token: string
    access_token_exp: number
    refresh_token: string
}

interface User {
    id?: string | null
    email?: string | null
    name?: string | null
    roles?: Role[] | null
    permissions?: Permission[] | null
}

type _AppUser = User

declare module "next-auth" {
    interface Session {
        user: User
        token: Token
    }

    interface User extends _AppUser {
        token?: Token
    }
}

declare module "next-auth/jwt" {
    interface JWT extends Token {
        user: _AppUser

        error?: "RefreshAccessTokenError"
    }
}