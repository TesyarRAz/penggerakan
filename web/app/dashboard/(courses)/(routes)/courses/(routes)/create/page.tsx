
import React from 'react'
import CreateCourseForm from './_components/create-course-form'
import { auth } from '@/auth'

const CreateCoursePage = async () => {
    const session = await auth()

    if (!session) return null

    return (
        <div>
            <h2>Create Course Page</h2>
            <CreateCourseForm session={session}/>
        </div>
    )
}

export default CreateCoursePage
