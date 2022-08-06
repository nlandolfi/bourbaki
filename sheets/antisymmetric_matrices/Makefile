# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make antisymmetric_matrices.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

antisymmetric_matrices.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > antisymmetric_matrices.tex

antisymmetric_matrices.pdf: ../../*.tex ../../trademark.pdf *.tex antisymmetric_matrices.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode antisymmetric_matrices.tex
	../../latexrun antisymmetric_matrices.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: antisymmetric_matrices.pdf
	open antisymmetric_matrices.pdf

o:
	make open

terms: antisymmetric_matrices.tex
	bbk_sheet -mode ts antisymmetric_matrices.tex

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
	rm -f antisymmetric_matrices.*
	make reset
	bbk_sheet -mode mk > Makefile