"use client"
import React, { Component, FunctionComponent, useState } from "react";
import {
    Dialog,
    DialogContent,
    DialogDescription, DialogFooter,
    DialogHeader,
    DialogTitle,
    DialogTrigger
} from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { FormItem, FormLabel } from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { SubModuleStructureType } from "@/types/enums";
import { Label } from "@/components/ui/label";
import StructureContent from "@/app/dashboard/(courses)/(routes)/_components/structure-content";
import { DialogProps } from "@radix-ui/react-dialog";

const PreviewStructureModal = ({
    structures,
    ...props
}: DialogProps & {
    structures: SubModuleStructure[],
}) => {
    return (
        <Dialog {...props}>
            <DialogContent>
                <DialogHeader>
                    <DialogTitle>Lihat Structure</DialogTitle>
                    <DialogDescription>
                        This action cannot be undone. This will permanently delete your account
                        and remove your data from our servers.
                    </DialogDescription>
                </DialogHeader>
                <div>
                    <StructureContent structures={structures} />
                </div>
            </DialogContent>
        </Dialog>
    )
}

export default PreviewStructureModal