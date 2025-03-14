import customtkinter as tk


class Skeleton(tk.CTk):

    def __init__(self, fg_color=None, **kwargs):
        super().__init__(fg_color, **kwargs)

        tk.CTkCheckBox(self, text="test").pack()

        self.mainloop()


Skeleton()
