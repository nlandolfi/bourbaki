# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make logistic_probabilistic_models.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

logistic_probabilistic_models.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > logistic_probabilistic_models.tex

logistic_probabilistic_models.pdf: ../../*.tex ../../trademark.pdf *.tex logistic_probabilistic_models.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode logistic_probabilistic_models.tex
	../../latexrun logistic_probabilistic_models.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: logistic_probabilistic_models.pdf
	open logistic_probabilistic_models.pdf

o:
	make open

terms: logistic_probabilistic_models.tex
	bbk_sheet -mode ts logistic_probabilistic_models.tex

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
	rm -f logistic_probabilistic_models.*
	make reset
	bbk_sheet -mode mk > Makefile