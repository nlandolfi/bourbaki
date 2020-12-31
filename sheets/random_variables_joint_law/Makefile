random_variables_joint_law.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex random_variables_joint_law.tex
	make defs

make:
	make random_variables_joint_law.pdf

open: random_variables_joint_law.pdf
	open random_variables_joint_law.pdf

o:
	make open

defs: random_variables_joint_law.tex
	defs random_variables_joint_law.tex
