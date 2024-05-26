"use client"
import React, { createContext, useEffect, useState } from 'react'
import { Session } from "next-auth";
import { Button } from "@/components/ui/button";
import { Popover, PopoverContent, PopoverTrigger } from '@/components/ui/popover';
import { Command, CommandEmpty, CommandGroup, CommandInput, CommandItem, CommandList } from '@/components/ui/command';
import StructureModal from './structure-modal';
import { SubModuleStructureType } from '@/types/enums';
import { v4 as uuid } from 'uuid';
import { Collapsible, CollapsibleContent, CollapsibleTrigger } from '@/components/ui/collapsible';
import { FaPlus, FaSort, FaSortDown, FaSortUp } from 'react-icons/fa';
import { cn } from '@/lib/utils';
import PreviewStructureModal from './preview-structure-modal';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { zodResolver } from '@hookform/resolvers/zod';
import { Label } from '@/components/ui/label';
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form';
import { Input } from '@/components/ui/input';
import editSubModule from '../_actions/edit-submodule-action';
import { useRouter } from 'next/navigation';
import createSubModule from '../_actions/create-submodule-action';
import Link from 'next/link';

const AddItemButton = ({
    label,
    addItem
}: {
    label?: React.ReactNode,
    addItem: (type: string) => void
}) => {
    const [open, setOpen] = useState(false)

    return (
        <Popover open={open} onOpenChange={setOpen}>
            <PopoverTrigger asChild>
                <Button type="button" variant="outline">
                    {label ?? "Add Item"}
                </Button>
            </PopoverTrigger>
            <PopoverContent className="p-0" side="right" align="start">
                <Command>
                    <CommandInput placeholder="Add item..." />
                    <CommandList>
                        <CommandEmpty>No results found.</CommandEmpty>
                        <CommandGroup>
                            {Object.values(SubModuleStructureType).map((value) => (
                                <CommandItem
                                    key={value}
                                    value={value}
                                    onSelect={(value) => {
                                        addItem(value)
                                        setOpen(false)
                                    }}
                                >
                                    {value}
                                </CommandItem>
                            ))}
                        </CommandGroup>
                    </CommandList>
                </Command>
            </PopoverContent>
        </Popover>
    )
}

const ListItem = ({
    structure,
    setSelectedStructure,
    updateItems,
    deleteItem,
    ...props
}: React.HTMLAttributes<HTMLDivElement> & {
    structure: SubModuleStructure,
    setSelectedStructure: (structure: SubModuleStructure) => void,
    updateItems: () => void,
    deleteItem: () => void,
}) => {
    const [open, setOpen] = useState(true)

    return (
        <Collapsible open={open} onOpenChange={setOpen}>
            <div className="ml-1" {...props}>
                <div className="flex justify-between">
                    <div>
                        <CollapsibleTrigger asChild>
                            <Button className={cn("mr-3", (structure.children?.length ?? 0) == 0 ? 'invisible' : '')} variant="ghost">
                                <FaSortUp className={cn("h-4 w-4", open ? "hidden" : "")} />
                                <FaSortDown className={cn("h-4 w-4", !open ? "hidden" : "")} />
                            </Button>
                        </CollapsibleTrigger>
                        {structure.label}
                    </div>
                    <div>
                        <AddItemButton
                            label={(
                                <>
                                    <FaPlus className="mr-1" />
                                </>
                            )}
                            addItem={(type) => {
                                structure.children = [
                                    ...structure.children ?? [],
                                    {
                                        resource_id: uuid(),
                                        resource_type: type,
                                        label: "New Item",
                                        value: ""
                                    }
                                ]

                                updateItems()
                            }} />
                        <Button
                            type="button"
                            variant="outline"
                            onClick={() => {
                                setSelectedStructure(structure)
                            }}
                        >
                            Edit
                        </Button>
                        <Button
                            type="button"
                            variant="outline"
                            onClick={deleteItem}
                        >
                            Delete
                        </Button>
                    </div>
                </div>

                <div className="mb-3">
                    <CollapsibleContent>
                        {structure.children?.map((item, index) => (
                            <ListItem
                                key={index}
                                structure={item}
                                updateItems={updateItems}
                                deleteItem={() => {
                                    structure.children = structure.children?.filter((_, i) => i !== index)
                                    updateItems()
                                }}
                                setSelectedStructure={setSelectedStructure} />
                        ))}
                    </CollapsibleContent>
                </div>
            </div>
        </Collapsible>
    )
}

const submoduleSchema = z.object({
    name: z.string().min(2, {
        message: "Name is too short",
    }),

})

const SubModuleForm = ({
    session,
    mod,
    submodule,
}: {
    session: Session,
    mod: ModuleResponse,
    submodule?: SubModuleResponse,
}) => {
    const router = useRouter()

    const [modalOpen, setModalOpen] = useState(false)
    const [structureModalOpen, setStructureModalOpen] = useState(false)

    const [selectedStructure, setSelectedStructure] = useState<SubModuleStructure | null>(null)
    const [items, setItems] = useState<SubModuleStructure[]>(submodule?.structure ?? [])

    const form = useForm<z.infer<typeof submoduleSchema>>({
        resolver: zodResolver(submoduleSchema),
        defaultValues: {
            name: submodule?.name ?? "",
        }
    })

    const { isSubmitting } = form.formState

    useEffect(() => {
        if (selectedStructure !== null) {
            setModalOpen(true)
        }
    }, [selectedStructure, setModalOpen])

    const addItem = (type: string) => {
        setItems([...items, {
            resource_id: uuid(),
            resource_type: type,
            label: "New Item",
            value: ""
        }])
    }

    const handleModalSubmit = (label: string, value: string) => {
        if (selectedStructure) {
            selectedStructure.label = label
            selectedStructure.value = value
        }
        setModalOpen(false)
        setSelectedStructure(null)
        setItems([...items])
    }

    const onSubmit = (values: z.infer<typeof submoduleSchema>) => {
        let response
        if (submodule) {
            response = editSubModule(session, mod.id, submodule.id, {
                name: values.name,
                structure: items
            })
        } else {
            response = createSubModule(session, mod.id, {
                name: values.name,
                structure: items
            })
        }

        alert('Berhasil mengedit submodule')

        router.refresh()
    }

    return (
        <>
            <StructureModal open={modalOpen} onOpenChange={(open) => {
                setModalOpen(open)
                if (!open) setSelectedStructure(null)
            }} structure={selectedStructure} onSubmit={handleModalSubmit} />
            <PreviewStructureModal open={structureModalOpen} onOpenChange={setStructureModalOpen} structures={items} />

            <Form {...form}>
                <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4 md:space-y-6">
                    <FormField
                        control={form.control}
                        name="name"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>
                                    Name
                                </FormLabel>
                                <FormControl>
                                    <Input
                                        disabled={isSubmitting}
                                        {...field}
                                    />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        )}
                    />
                    <div className="border rounded-lg p-3">

                        <div className="flex justify-between">
                            <h4>Structure</h4>
                            <Button type="button" onClick={() => setStructureModalOpen(true)}>Preview</Button>
                        </div>
                        <div className="mt-5 ">
                            {items.map((item, index) => (
                                <ListItem
                                    key={index}
                                    structure={item}
                                    updateItems={() => setItems([...items])}
                                    deleteItem={() => setItems(items.filter((_, i) => i !== index))}
                                    setSelectedStructure={setSelectedStructure} />
                            ))}

                            <div className="mt-5">
                                <AddItemButton addItem={addItem} />
                            </div>
                        </div>
                    </div>

                    <Button type="submit" disabled={isSubmitting}>Save</Button>
                </form>
            </Form>
        </>
    )
}

export default SubModuleForm
