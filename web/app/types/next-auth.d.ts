import "next-auth"
import "next-auth/jwt"

interface Token {
    access_token: string
    access_token_expires: number
    refresh_token: string
}

declare module "next-auth" {
    interface Session {
        error?: "RefreshAccessTokenError"

        token: Token
    }

    interface User {
        id: string
        email: string
        name: string
        roles: Role[]
        permissions: Permission[]
    }
}

declare module "next-auth/jwt" {
    interface JWT extends Token {
        error?: "RefreshAccessTokenError"

        user: User
        token: Token
    }
}