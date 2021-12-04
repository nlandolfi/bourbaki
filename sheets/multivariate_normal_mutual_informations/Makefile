# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make multivariate_normal_mutual_informations.pdf

multivariate_normal_mutual_informations.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > multivariate_normal_mutual_informations.tex

multivariate_normal_mutual_informations.pdf: ../../*.tex ../../trademark.pdf *.tex multivariate_normal_mutual_informations.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode multivariate_normal_mutual_informations.tex
	../../latexrun multivariate_normal_mutual_informations.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: multivariate_normal_mutual_informations.pdf
	open multivariate_normal_mutual_informations.pdf

o:
	make open

terms: multivariate_normal_mutual_informations.tex
	bbk_sheet -mode ts multivariate_normal_mutual_informations.tex

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
	rm -f multivariate_normal_mutual_informations.*
	make reset
	bbk_sheet -mode mk > Makefile