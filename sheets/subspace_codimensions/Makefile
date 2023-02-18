# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make subspace_codimensions.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

subspace_codimensions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > subspace_codimensions.tex

subspace_codimensions.pdf: ../../*.tex ../../trademark.pdf *.tex subspace_codimensions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode subspace_codimensions.tex
	../../latexrun subspace_codimensions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: subspace_codimensions.pdf
	open subspace_codimensions.pdf

o:
	make open

terms: subspace_codimensions.tex
	bbk_sheet -mode ts subspace_codimensions.tex

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
	rm -f subspace_codimensions.*
	make reset
	bbk_sheet -mode mk > Makefile