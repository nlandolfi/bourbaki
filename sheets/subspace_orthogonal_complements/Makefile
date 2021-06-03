# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make subspace_orthogonal_complements.pdf

subspace_orthogonal_complements.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > subspace_orthogonal_complements.tex

subspace_orthogonal_complements.pdf: ../../*.tex ../../trademark.pdf *.tex subspace_orthogonal_complements.tex
	# pdflatex --file-line-error -interaction=nonstopmode subspace_orthogonal_complements.tex
	../../latexrun subspace_orthogonal_complements.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: subspace_orthogonal_complements.pdf
	open subspace_orthogonal_complements.pdf

o:
	make open

terms: subspace_orthogonal_complements.tex
	bbk_sheet -mode ts subspace_orthogonal_complements.tex

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
	rm -f subspace_orthogonal_complements.*
	make reset
	bbk_sheet -mode mk > Makefile