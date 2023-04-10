# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make probabilistic_dataset_models.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

probabilistic_dataset_models.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > probabilistic_dataset_models.tex

probabilistic_dataset_models.pdf: ../../*.tex ../../trademark.pdf *.tex probabilistic_dataset_models.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode probabilistic_dataset_models.tex
	../../latexrun probabilistic_dataset_models.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: probabilistic_dataset_models.pdf
	open probabilistic_dataset_models.pdf

o:
	make open

terms: probabilistic_dataset_models.tex
	bbk_sheet -mode ts probabilistic_dataset_models.tex

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
	rm -f probabilistic_dataset_models.*
	make reset
	bbk_sheet -mode mk > Makefile