import { Form } from '@/components/ui/form'
import { moduleSchema } from '@/lib/zod'
import { zodResolver } from '@hookform/resolvers/zod'
import { Session } from 'next-auth'
import { useSession } from 'next-auth/react'
import { useRouter } from 'next/navigation'
import React from 'react'
import { useForm } from 'react-hook-form'
import { z } from 'zod'
import NameForm from '../../_components/name-form'
import { Button } from '@/components/ui/button'

const CreateModuleForm = async ({
    course
}: {
    course: CourseResponse
}) => {
    const { data: session, status } = useSession()
    const router = useRouter()

    if (status === "unauthenticated")
        router.push(`/auth/signin?callback=/dashboard/modules/${course.id}/create`)

    const form = useForm<z.infer<typeof moduleSchema>>({
        resolver: zodResolver(moduleSchema),
        defaultValues: {
            name: "",
        },
    });

    const { isSubmitting } = form.formState

    const onSubmit = async (values: z.infer<typeof moduleSchema>) => {

    }

    return (
        <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4 md:space-y-6">
                <NameForm control={form.control} isSubmitting={isSubmitting} />
                <Button type="submit" disabled={isSubmitting}>
                    Simpan
                </Button>
            </form>
        </Form>
    )
}

export default CreateModuleForm
