# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make joint_distributions.pdf
	make graph.pdf

joint_distributions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > joint_distributions.tex

joint_distributions.pdf: ../../*.tex ../../trademark.pdf *.tex joint_distributions.tex
	pdflatex --file-line-error -interaction=nonstopmode joint_distributions.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: joint_distributions.pdf
	open joint_distributions.pdf

o:
	make open

defs: joint_distributions.tex
	defs joint_distributions.tex

spell:
	aspell -c joint_distributions.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm joint_distributions.*
	make reset
	bbk_sheet -mode mk > Makefile