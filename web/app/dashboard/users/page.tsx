import Link from 'next/link'
import React from 'react'
import { FaPlus, FaSearch } from 'react-icons/fa'
import { RiEqualizerFill } from 'react-icons/ri'
import CourseCard from '../(courses)/courses/_components/course-card'
import getUsers from './_actions/get-users-action'
import { getServerSession } from 'next-auth'
import { authOptions } from '@/app/api/auth/[...nextauth]/route'
import { redirect } from 'next/navigation'
import UserCard from './_components/user-card'
import { IoAddCircleOutline } from 'react-icons/io5'
import { Button } from '@/components/ui/button'
import BrowseLayout from '@/components/browse-layout'

const UsersPage = async () => {
  const session = await getServerSession(authOptions)

  if (!session) {
    return redirect("/auth/signin?callback=/dashboard/users");
  }

  const users = await getUsers(session)
  
  return (
    <BrowseLayout
      title="User List"
      tools={(
        <Button
          variant="blank"
          className="text-green-700 hover:text-white border border-green-700 hover:bg-green-800 focus:ring-4 focus:outline-none focus:ring-green-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center me-2 mb-2 dark:border-green-500 dark:text-green-500 dark:hover:text-white dark:hover:bg-green-600 dark:focus:ring-green-800"
          asChild
        >
          <Link href="/dashboard/users/create">
            <IoAddCircleOutline className="mr-1 h-5 w-5" />
            <span>Add User</span>
          </Link>
        </Button >
      )}
    >
      <div className="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 mt-5">
        {users.data.map((item) => (
          <UserCard
            key={item.id}
            user={item}
          />
        ))}
      </div>
    </BrowseLayout>
  )
}

export default UsersPage
