# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make vectors_as_matrices.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

vectors_as_matrices.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > vectors_as_matrices.tex

vectors_as_matrices.pdf: ../../*.tex ../../trademark.pdf *.tex vectors_as_matrices.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode vectors_as_matrices.tex
	../../latexrun vectors_as_matrices.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: vectors_as_matrices.pdf
	open vectors_as_matrices.pdf

o:
	make open

terms: vectors_as_matrices.tex
	bbk_sheet -mode ts vectors_as_matrices.tex

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
	rm -f vectors_as_matrices.*
	make reset
	bbk_sheet -mode mk > Makefile