import { auth } from '@/auth';
import { redirect } from 'next/navigation';
import React from 'react'
import getTeachers from '../../../_actions/get-teachers-action';
import BrowseLayout from '@/components/layouts/browse-layout';
import { Button } from '@/components/ui/button';
import Link from 'next/link';
import { IoAddCircleOutline } from 'react-icons/io5';
import TeacherCard from './_components/teacher-card';

const TeacherPage = async () => {
  const session = await auth()

  if (!session) {
    return redirect("/auth/signin?callback=/dashboard/teachers");
  }

  const teachers = await getTeachers(session)

  return (
    <BrowseLayout
      title="Teacher List"
      tools={(
        <Button
          variant="blank"
          className="text-green-700 hover:text-white border border-green-700 hover:bg-green-800 focus:ring-4 focus:outline-none focus:ring-green-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center me-2 mb-2 dark:border-green-500 dark:text-green-500 dark:hover:text-white dark:hover:bg-green-600 dark:focus:ring-green-800"
          asChild
        >
          <Link href="/dashboard/teachers/create">
            <IoAddCircleOutline className="mr-1 h-5 w-5" />
            <span>Add Teacher</span>
          </Link>
        </Button >
      )}
    >
      <div className="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 mt-5">
        {teachers.data.map((item) => (
          <TeacherCard
            session={session}
            key={item.id}
            teacher={item}
          />
        ))}
      </div>
    </BrowseLayout>
  )
}

export default TeacherPage
