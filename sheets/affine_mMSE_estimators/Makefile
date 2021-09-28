# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make affine_mMSE_estimators.pdf

affine_mMSE_estimators.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > affine_mMSE_estimators.tex

affine_mMSE_estimators.pdf: ../../*.tex ../../trademark.pdf *.tex affine_mMSE_estimators.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode affine_mMSE_estimators.tex
	../../latexrun affine_mMSE_estimators.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: affine_mMSE_estimators.pdf
	open affine_mMSE_estimators.pdf

o:
	make open

terms: affine_mMSE_estimators.tex
	bbk_sheet -mode ts affine_mMSE_estimators.tex

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
	rm -f affine_mMSE_estimators.*
	make reset
	bbk_sheet -mode mk > Makefile