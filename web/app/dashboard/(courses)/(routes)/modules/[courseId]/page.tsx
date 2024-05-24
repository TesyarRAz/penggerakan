import React from 'react'
import { notFound, redirect } from 'next/navigation'
import { Button } from '@/components/ui/button'
import BrowseLayout from '@/components/layouts/browse-layout'
import Link from 'next/link'
import { IoAddCircleOutline } from 'react-icons/io5'
import { auth } from '@/auth'
import getCourseById from '@/app/dashboard/(courses)/_actions/get-course-byid-action'
import getModules from './_actions/get-modules-action'
import ModuleCard from './_components/module-card'

const ModulesPage = async ({
    params
}: {
    params: {
        courseId: string
    }
}) => {
    const session = await auth()

    if (!session) {
        return redirect(`/auth/signin?callback=/dashboard/modules/${params.courseId}`)
    }

    const course = await getCourseById(session, params.courseId)
    const modules = await getModules(session, params.courseId)

    return (
        <BrowseLayout
            title="Module List"
            tools={(
                <Button
                    variant="blank"
                    className="text-green-700 hover:text-white border border-green-700 hover:bg-green-800 focus:ring-4 focus:outline-none focus:ring-green-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center me-2 mb-2 dark:border-green-500 dark:text-green-500 dark:hover:text-white dark:hover:bg-green-600 dark:focus:ring-green-800"
                    asChild
                >
                    <Link href={`/dashboard/modules/${course.id}/create`}>
                        <IoAddCircleOutline className="mr-1 h-5 w-5" />
                        <span>Add Module</span>
                    </Link>
                </Button>
            )}
        >
            <div className="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 mt-5">
                {modules.data.map((item) => (
                    <ModuleCard
                        session={session}
                        key={item.id}
                        course={course}
                        module={item}
                    />
                ))}
            </div>
        </BrowseLayout>
    )
}

export default ModulesPage
