# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make trees.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

trees.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > trees.tex

trees.pdf: ../../*.tex ../../trademark.pdf *.tex trees.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode trees.tex
	../../latexrun trees.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: trees.pdf
	open trees.pdf

o:
	make open

terms: trees.tex
	bbk_sheet -mode ts trees.tex

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
	rm -f trees.*
	make reset
	bbk_sheet -mode mk > Makefile