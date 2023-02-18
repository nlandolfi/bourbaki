# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make subspace_dimensions.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

subspace_dimensions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > subspace_dimensions.tex

subspace_dimensions.pdf: ../../*.tex ../../trademark.pdf *.tex subspace_dimensions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode subspace_dimensions.tex
	../../latexrun subspace_dimensions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: subspace_dimensions.pdf
	open subspace_dimensions.pdf

o:
	make open

terms: subspace_dimensions.tex
	bbk_sheet -mode ts subspace_dimensions.tex

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
	rm -f subspace_dimensions.*
	make reset
	bbk_sheet -mode mk > Makefile