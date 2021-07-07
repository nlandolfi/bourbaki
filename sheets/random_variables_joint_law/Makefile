# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make random_variables_joint_law.pdf

random_variables_joint_law.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > random_variables_joint_law.tex

random_variables_joint_law.pdf: ../../*.tex ../../trademark.pdf *.tex random_variables_joint_law.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode random_variables_joint_law.tex
	../../latexrun random_variables_joint_law.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: random_variables_joint_law.pdf
	open random_variables_joint_law.pdf

o:
	make open

terms: random_variables_joint_law.tex
	bbk_sheet -mode ts random_variables_joint_law.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f random_variables_joint_law.*
	make reset
	bbk_sheet -mode mk > Makefile