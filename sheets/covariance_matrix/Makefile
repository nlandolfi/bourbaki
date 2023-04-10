# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make covariance_matrix.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

covariance_matrix.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > covariance_matrix.tex

covariance_matrix.pdf: ../../*.tex ../../trademark.pdf *.tex covariance_matrix.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode covariance_matrix.tex
	../../latexrun covariance_matrix.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: covariance_matrix.pdf
	open covariance_matrix.pdf

o:
	make open

terms: covariance_matrix.tex
	bbk_sheet -mode ts covariance_matrix.tex

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
	rm -f covariance_matrix.*
	make reset
	bbk_sheet -mode mk > Makefile