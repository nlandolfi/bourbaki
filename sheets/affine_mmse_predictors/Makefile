# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make affine_mmse_predictors.pdf
	make graph.pdf

affine_mmse_predictors.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > affine_mmse_predictors.tex

affine_mmse_predictors.pdf: ../../*.tex ../../trademark.pdf *.tex affine_mmse_predictors.tex
	# pdflatex --file-line-error -interaction=nonstopmode affine_mmse_predictors.tex
	../../latexrun affine_mmse_predictors.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: affine_mmse_predictors.pdf
	open affine_mmse_predictors.pdf

o:
	make open

terms: affine_mmse_predictors.tex
	bbk_sheet -mode ts affine_mmse_predictors.tex

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
	rm -f affine_mmse_predictors.*
	make reset
	bbk_sheet -mode mk > Makefile