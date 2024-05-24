import getTeachers from '@/app/dashboard/_actions/get-teachers-action'
import { Button } from '@/components/ui/button'
import { Command, CommandEmpty, CommandGroup, CommandInput, CommandItem, CommandList } from '@/components/ui/command'
import { FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form'
import { Popover, PopoverContent, PopoverTrigger } from '@/components/ui/popover'
import { cn } from '@/lib/utils'
import { Session } from 'next-auth'
import React, { useDeferredValue, useEffect, useOptimistic, useState } from 'react'
import {useFormContext, UseFormReturn} from 'react-hook-form'
import { FaSort } from 'react-icons/fa'
import { RiCheckLine } from 'react-icons/ri'

const TeacherForm = ({
    session,
}: {
    session: Session,
}) => {
    const form = useFormContext()

    const { isSubmitting } = form.formState

    const [search, setSearch] = useState<string>("")
    const deferredSearch = useDeferredValue(search)

    const [teachers, setTeachers] = useState<TeacherResponse[]>([])

    useEffect(() => {
        const fetchTeachers = setTimeout(async () => {
            const teachers = await getTeachers(session, {
                search: deferredSearch
            })

            setTeachers(teachers.data)
        }, 1000)

        return () => {
            clearTimeout(fetchTeachers)
        }
    }, [deferredSearch, session])

    useEffect(() => {
        console.log(teachers.map(teacher => teacher.name))
    }, [teachers])

    return (
        <FormField
            control={form.control}
            name="teacher_id"
            render={({ field }) => (
                <FormItem className="flex flex-col">
                    <FormLabel>
                        Teacher ID
                    </FormLabel>
                    <Popover>
                        <PopoverTrigger asChild>
                            <FormControl>
                                <Button
                                    variant="outline"
                                    role="combobox"
                                    className={cn(
                                        "w-[200px] justify-between",
                                        !field.value && "text-muted-foreground"
                                    )}
                                >
                                    {field.value
                                        ? teachers.find(
                                            (teacher) => teacher.id === field.value
                                        )?.name
                                        : "Pilih Guru"}
                                    <FaSort className="ml-2 h-4 w-4 shrink-0 opacity-50" />
                                </Button>
                            </FormControl>
                        </PopoverTrigger>
                        <PopoverContent className="w-[200px] p-0">
                            <Command shouldFilter={false}>
                                <CommandInput
                                    disabled={isSubmitting}
                                    placeholder="Cari guru..."
                                    className="h-9"
                                    onInput={(e) => setSearch(e.currentTarget.value)}
                                />
                                <CommandList>
                                    <CommandEmpty>Guru tidak ditemukan</CommandEmpty>
                                    <CommandGroup>
                                        {teachers.map((teacher) => (
                                            <CommandItem
                                                value={teacher.id}
                                                key={teacher.id}
                                                onSelect={() => {
                                                    form.setValue("teacher_id", field.value != teacher.id ? teacher.id : null)
                                                }}
                                            >
                                                {teacher.name}
                                                <RiCheckLine
                                                    className={cn(
                                                        "ml-auto h-4 w-4",
                                                        teacher.id === field.value
                                                            ? "opacity-100"
                                                            : "opacity-0"
                                                    )}
                                                />
                                            </CommandItem>
                                        ))}
                                    </CommandGroup>
                                </CommandList>
                            </Command>
                        </PopoverContent>
                    </Popover>
                    <FormDescription>
                        Guru yang akan mengurus course
                    </FormDescription>
                    <FormMessage />
                </FormItem>
            )}
        />
    )
}

export default TeacherForm
