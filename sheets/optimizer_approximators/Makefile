# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make optimizer_approximators.pdf

optimizer_approximators.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > optimizer_approximators.tex

optimizer_approximators.pdf: ../../*.tex ../../trademark.pdf *.tex optimizer_approximators.tex
	# pdflatex --file-line-error -interaction=nonstopmode optimizer_approximators.tex
	../../latexrun optimizer_approximators.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: optimizer_approximators.pdf
	open optimizer_approximators.pdf

o:
	make open

terms: optimizer_approximators.tex
	bbk_sheet -mode ts optimizer_approximators.tex

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
	rm -f optimizer_approximators.*
	make reset
	bbk_sheet -mode mk > Makefile