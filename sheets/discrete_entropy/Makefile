# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make discrete_entropy.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

discrete_entropy.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > discrete_entropy.tex

discrete_entropy.pdf: ../../*.tex ../../trademark.pdf *.tex discrete_entropy.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode discrete_entropy.tex
	../../latexrun discrete_entropy.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: discrete_entropy.pdf
	open discrete_entropy.pdf

o:
	make open

terms: discrete_entropy.tex
	bbk_sheet -mode ts discrete_entropy.tex

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
	rm -f discrete_entropy.*
	make reset
	bbk_sheet -mode mk > Makefile